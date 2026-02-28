<script lang="ts">
	import { SquareTerminal } from 'lucide-svelte';
</script>

<div
	class="flex h-7 items-center justify-between border-t border-sidebar-border bg-card px-3 text-[11px] text-muted-foreground"
>
	<div class="flex items-center gap-3">
		<div
			class="flex items-center gap-1 rounded px-1.5 py-0.5 text-yellow-500 bg-yellow-500/10"
		>
			<span>Free</span>
		</div>
		<span>Left Only 30 day</span>
		<button
			class="rounded bg-primary px-2.5 py-0.5 text-[11px] font-medium text-primary-foreground hover:bg-primary/90"
		>
			Upgrade
		</button>
	</div>
	<div class="flex items-center gap-2">
		<button class="flex items-center gap-1 hover:text-foreground">
			<SquareTerminal size={12} />
		</button>
	</div>
</div>
