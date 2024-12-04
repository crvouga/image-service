import { $computed, $signal } from "../ui/signals.js";

export const $document = $signal(document);

let mutationTimeout;
const observer = new MutationObserver(() => {
  clearTimeout(mutationTimeout);
  mutationTimeout = setTimeout(() => $document.set(document), 50);
});

observer.observe(document.documentElement, {
  childList: true,
  subtree: true,
});

/**
 * Reactive querySelector
 * @param {import("./signals").Computed<Element | Document>} $node
 * @param {string} selector
 * @returns {import("./signals").Computed<HTMLElement | null>}
 */
export const $querySelector = ($node, selector) => {
  return $computed(() => {
    try {
      const found = $node.get()?.querySelector(selector);
      return found instanceof HTMLElement ? found : null;
    } catch (e) {
      console.error("Error in $querySelector:", e);
      return null;
    }
  });
};

/**
 * Reactive closest selector
 * @param {import("./signals").Computed<Element | Document>} $node
 * @param {string} selector
 * @returns {import("./signals").Computed<HTMLElement | null>}
 */
export const $closest = ($node, selector) => {
  return $computed(() => {
    try {
      const found = $node.get();
      return found instanceof HTMLElement ? found.closest(selector) : null;
    } catch (e) {
      console.error("Error in $closest:", e);
      return null;
    }
  });
};

// Ensure MutationObserver is disconnected if needed (e.g., app teardown)
export const disconnectObserver = () => {
  observer.disconnect();
};
