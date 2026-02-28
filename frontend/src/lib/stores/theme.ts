import { writable } from 'svelte/store';
import { browser } from '$app/environment';

type Theme = 'light' | 'dark';

function createThemeStore() {
	const initial: Theme = browser
		? (localStorage.getItem('theme') as Theme) ??
			(window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light')
		: 'light';

	const { subscribe, set, update } = writable<Theme>(initial);

	function apply(theme: Theme) {
		if (browser) {
			document.documentElement.classList.toggle('dark', theme === 'dark');
			localStorage.setItem('theme', theme);
		}
	}

	// Apply initial theme immediately
	apply(initial);

	return {
		subscribe,
		toggle() {
			update((current) => {
				const next: Theme = current === 'dark' ? 'light' : 'dark';
				apply(next);
				return next;
			});
		},
		set(theme: Theme) {
			apply(theme);
			set(theme);
		}
	};
}

export const theme = createThemeStore();
