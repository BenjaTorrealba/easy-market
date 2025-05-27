// src/lib/localStorageStore.ts
// 
import { writable, type Writable } from 'svelte/store';

export function localStorageStore<T>(key: string, initialValue: T): Writable<T> {
  let storedValue: T;

  try {
    const json = localStorage.getItem(key);
    storedValue = json ? JSON.parse(json) as T : initialValue;
  } catch (e) {
    console.warn(`Error reading localStorage key "${key}":`, e);
    storedValue = initialValue;
  }

  const store = writable<T>(storedValue);

  store.subscribe((value) => {
    try {
      localStorage.setItem(key, JSON.stringify(value));
    } catch (e) {
      console.warn(`Error writing to localStorage key "${key}":`, e);
    }
  });

  return store;
}
