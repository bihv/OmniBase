<script lang="ts">
	import { ChevronRight, ChevronDown, Folder } from 'lucide-svelte';
	import { connectionsStore, type Group } from '$lib/stores/connections';

	export let group: Group;

	function autoFocus(node: HTMLInputElement) {
		node.focus();
		node.select();
		return {};
	}

	function handleRename(e: Event) {
		const value = (e.currentTarget as HTMLInputElement).value;
		connectionsStore.renameGroup(group.id, value);
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') {
			connectionsStore.renameGroup(group.id, (e.currentTarget as HTMLInputElement).value);
		} else if (e.key === 'Escape') {
			connectionsStore.cancelEdit(group.id);
		}
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	class="flex cursor-pointer select-none items-center gap-1 px-2 py-1 text-foreground hover:bg-accent"
	on:click={() => !group.editing && connectionsStore.toggleExpand(group.id)}
>
	<span class="flex h-4 w-4 items-center justify-center text-muted-foreground">
		{#if group.expanded}
			<ChevronDown size={12} />
		{:else}
			<ChevronRight size={12} />
		{/if}
	</span>

	<Folder size={15} class="shrink-0 text-amber-400" />

	{#if group.editing}
		<!-- svelte-ignore a11y-click-events-have-key-events -->
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		<input
			use:autoFocus
			type="text"
			value={group.name}
			class="ml-1 flex-1 rounded border border-ring bg-ring/20 px-1.5 py-0.5 text-sm text-foreground outline-none"
			on:click|stopPropagation
			on:blur={handleRename}
			on:keydown={handleKeydown}
		/>
	{:else}
		<span class="ml-1 flex-1 truncate text-sm">{group.name}</span>
	{/if}
</div>
