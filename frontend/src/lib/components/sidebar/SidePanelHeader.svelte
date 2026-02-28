<script lang="ts">
	import { Globe, EllipsisVertical } from 'lucide-svelte';
</script>

<div class="flex items-center justify-between px-3 py-2.5">
	<span class="text-sm font-medium text-foreground">Database</span>
	<div class="flex items-center gap-0.5">
		<button
			class="flex h-6 w-6 items-center justify-center rounded text-muted-foreground hover:bg-accent hover:text-foreground"
			title="Connections"
		>
			<Globe size={14} />
		</button>
		<button
			class="flex h-6 w-6 items-center justify-center rounded text-muted-foreground hover:bg-accent hover:text-foreground"
			title="More"
		>
			<EllipsisVertical size={14} />
		</button>
	</div>
</div>
