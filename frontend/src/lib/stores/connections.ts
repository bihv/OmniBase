import { writable } from 'svelte/store';

export type Group = {
	id: number;
	name: string;
	expanded: boolean;
	editing: boolean;
};

const STORAGE_KEY = 'omnibase:connections';

function loadFromStorage(): Group[] {
	try {
		const raw = localStorage.getItem(STORAGE_KEY);
		if (!raw) return [];
		const parsed: Group[] = JSON.parse(raw);
		// Reset editing state on load
		return parsed.map((g) => ({ ...g, editing: false }));
	} catch {
		return [];
	}
}

function saveToStorage(groups: Group[]) {
	try {
		localStorage.setItem(STORAGE_KEY, JSON.stringify(groups));
	} catch {
		// storage may be unavailable
	}
}

function createConnectionsStore() {
	const initial = loadFromStorage();
	let nextId = initial.length > 0 ? Math.max(...initial.map((g) => g.id)) + 1 : 1;

	const { subscribe, update } = writable<Group[]>(initial);

	function commit(updater: (groups: Group[]) => Group[]) {
		update((groups) => {
			const next = updater(groups);
			saveToStorage(next);
			return next;
		});
	}

	return {
		subscribe,

		addGroup() {
			const newGroup: Group = {
				id: nextId++,
				name: 'New Group',
				expanded: false,
				editing: true
			};
			commit((groups) => [...groups, newGroup]);
		},

		renameGroup(id: number, name: string) {
			commit((groups) =>
				groups.map((g) =>
					g.id === id ? { ...g, name: name.trim() || 'New Group', editing: false } : g
				)
			);
		},

		cancelEdit(id: number) {
			commit((groups) => groups.map((g) => (g.id === id ? { ...g, editing: false } : g)));
		},

		toggleExpand(id: number) {
			commit((groups) =>
				groups.map((g) => (g.id === id ? { ...g, expanded: !g.expanded } : g))
			);
		},

		deleteGroup(id: number) {
			commit((groups) => groups.filter((g) => g.id !== id));
		}
	};
}

export const connectionsStore = createConnectionsStore();
