<script lang="ts">
	import { Info, Code2, Copy, Sparkles } from 'lucide-svelte';

	const items = [
		{ icon: Info, label: 'Info' },
		{ icon: Code2, label: 'Schema' },
		{ icon: Copy, label: 'Snippets' },
		{ icon: Sparkles, label: 'AI' }
	];
</script>

<div class="flex h-full w-10 flex-col items-center gap-1 border-l border-sidebar-border bg-card py-2">
	{#each items as item}
		<button
			class="flex h-8 w-8 items-center justify-center rounded text-muted-foreground transition-colors hover:bg-accent hover:text-foreground"
			title={item.label}
		>
			<item.icon size={16} />
		</button>
	{/each}
</div>
