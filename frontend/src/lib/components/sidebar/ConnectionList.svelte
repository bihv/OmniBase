<script lang="ts">
	import { PackageOpen } from 'lucide-svelte';
	import { connectionsStore } from '$lib/stores/connections';
	import GroupRow from './GroupRow.svelte';
</script>

<div class="flex flex-1 flex-col overflow-y-auto">
	{#if $connectionsStore.length === 0}
		<div class="flex flex-1 flex-col items-center justify-center px-4">
			<div class="flex flex-col items-center gap-3 text-center">
				<div class="text-muted-foreground/60">
					<PackageOpen size={64} strokeWidth={1} />
				</div>
				<p class="text-xs text-muted-foreground">You have not created a connection yet</p>
			</div>
		</div>
	{:else}
		{#each $connectionsStore as group (group.id)}
			<GroupRow {group} />
		{/each}
	{/if}
</div>
