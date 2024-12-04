document.addEventListener("DOMContentLoaded", () => {
  const input = document.querySelector("input[name=email]");
  const inputHelper = input.closest("label").querySelector("small[role=alert]");
  const form = document.querySelector("form");
  const submitButton = document.querySelector("button[type=submit]");

  if (
    !(
      input instanceof HTMLInputElement &&
      inputHelper instanceof HTMLElement &&
      form instanceof HTMLFormElement &&
      submitButton instanceof HTMLButtonElement
    )
  ) {
    return;
  }

  if (inputHelper.textContent.length > 0) {
    input.setAttribute("aria-invalid", "true");
    inputHelper.id = "helper";
    input.setAttribute("aria-describedby", inputHelper.id);
    input.focus();
    input.setSelectionRange(input.value.length, input.value.length);
  }

  input.addEventListener("input", (e) => {
    input.removeAttribute("aria-invalid");
    input.removeAttribute("aria-describedby");
    const emailHelper = input
      .closest("label")
      .querySelector("small[role=alert]");

    emailHelper.textContent = "";
  });

  form.addEventListener("submit", (e) => {
    submitButton.setAttribute("aria-busy", "true");
  });
});
