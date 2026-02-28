<script lang="ts">
	import { Command } from 'lucide-svelte';
</script>

<div class="flex flex-1 flex-col items-center justify-center bg-background">
	<!-- App title -->
	<h1 class="mb-12 text-6xl font-light tracking-wider text-muted-foreground/50">OmniBase</h1>

	<!-- Keyboard shortcuts -->
	<div class="flex flex-col gap-4 text-sm">
		<div class="flex items-center justify-between gap-8">
			<span class="text-muted-foreground">Execute selected SQL</span>
			<div class="flex items-center gap-1">
				<kbd class="text-ring"><Command size={14} class="inline" /></kbd>
				<span class="text-muted-foreground/50">+</span>
				<kbd class="text-ring">R</kbd>
			</div>
		</div>

		<div class="flex items-center justify-between gap-8">
			<span class="text-muted-foreground">Save console</span>
			<div class="flex items-center gap-1">
				<kbd class="text-ring"><Command size={14} class="inline" /></kbd>
				<span class="text-muted-foreground/50">+</span>
				<kbd class="text-ring">S</kbd>
			</div>
		</div>

		<div class="flex items-center justify-between gap-8">
			<span class="text-muted-foreground">Create Console</span>
			<div class="flex items-center gap-1">
				<kbd class="text-ring"><Command size={14} class="inline" /></kbd>
				<span class="text-muted-foreground/50">+</span>
				<kbd class="text-ring">⇧</kbd>
				<span class="text-muted-foreground/50">+</span>
				<kbd class="text-ring">L</kbd>
			</div>
		</div>

		<div class="flex items-center justify-between gap-8">
			<span class="text-muted-foreground">Natural Language to SQL</span>
			<div class="flex items-center gap-1">
				<kbd class="text-ring">/</kbd>
			</div>
		</div>

		<div class="flex items-center justify-between gap-8">
			<span class="text-muted-foreground">Explain/Optimize SQL</span>
			<span class="text-ring text-xs">Right-Click of mouse</span>
		</div>
	</div>
</div>
