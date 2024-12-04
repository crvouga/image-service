// @ts-check
import { $closest, $document, $querySelector } from "../../ui/document.js";
import { $signal, effect, $computed } from "../../ui/signals.js";

const $emailHelperText = $signal("");
const $hasError = $computed(() => $emailHelperText.get().length > 0);
const $isSubmitting = $signal(false);

const $refForm = $querySelector($document, "form");
const $refSubmitButton = $querySelector($document, "button[type=submit]");
const $refEmail = $querySelector($document, "input[name=email]");
const $refEmailHelper = $computed(() =>
  $querySelector($closest($refEmail, "label"), "small[role=alert]").get()
);

effect(() => {
  $refForm.get().addEventListener("submit", () => $isSubmitting.set(true), {
    once: true,
  });
});

effect(() => {
  $emailHelperText.set($refEmailHelper.get().textContent || "");
  $refEmail.get().addEventListener("input", () => $emailHelperText.set(""), {
    once: true,
  });
});

effect(() => {
  if ($hasError.get()) {
    $refEmail.get().focus();
    $refEmail.get().setAttribute("aria-invalid", "true");
    $refEmail.get().setAttribute("aria-describedby", $refEmailHelper.get().id);
    $refEmailHelper.get().style.display = "block";
    $refEmailHelper.get().textContent = $emailHelperText.get();
  } else {
    $refEmail.get().removeAttribute("aria-invalid");
    $refEmail.get().removeAttribute("aria-describedby");
    $refEmailHelper.get().style.display = "none";
    $refEmailHelper.get().textContent = "";
  }
});

effect(() => {
  if ($isSubmitting.get()) {
    $refSubmitButton.get().setAttribute("aria-busy", "true");
  } else {
    $refSubmitButton.get().removeAttribute("aria-busy");
  }
});
