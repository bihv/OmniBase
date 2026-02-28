import { writable } from 'svelte/store';

export const connectionDialog = writable<{
	open: boolean;
	driver: string;
	editId: string | null;
}>({ open: false, driver: '', editId: null });
