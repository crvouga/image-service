/**
 * @type {() => void}
 */
let currentEffect = () => {};

/**
 * @template T
 * @typedef {{
 * get: () => T,
 * set: (newValue: T) => void
 * }} Signal
 */

/**
 * @template T
 * @param {T} initialValue
 * @returns {Signal<T>}
 */
export function $signal(initialValue) {
  let value = initialValue;
  /** @type {Set<() => void>} */
  const effects = new Set();
  return {
    /**
     * Retrieves the current value and registers the effect.
     * @returns {T}
     */
    get() {
      effects.add(currentEffect);
      return value;
    },
    /**
     * Updates the value and triggers effects if the value has changed.
     * @param {T} newValue
     * @returns {void}
     */
    set(newValue) {
      if (value !== newValue) {
        value = newValue;
        effects.forEach((callback) => callback());
      }
    },
  };
}

/**
 * Registers and executes a reactive effect.
 * @param {() => void} callback
 * @returns {void}
 */
export function effect(callback) {
  function wrappedEffect() {
    currentEffect = wrappedEffect;
    try {
      callback();
    } finally {
      currentEffect = () => {};
    }
  }
  wrappedEffect();
}

/**
 * @template T
 * @typedef {{
 * get: () => T,
 * }} Computed
 */

/**
 * Creates a computed signal that derives its value from other signals.
 * @template T
 * @param {() => T} computeFn - The function to compute the derived value.
 * @returns {Computed<T>}
 */
export function $computed(computeFn) {
  let cachedValue;
  const effects = new Set();

  effect(() => {
    cachedValue = computeFn();
    effects.forEach((callback) => callback());
  });

  return {
    /**
     * Retrieves the current value of the computed signal.
     * @returns {T}
     */
    get() {
      effects.add(currentEffect);
      return cachedValue;
    },
  };
}
