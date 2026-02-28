<script lang="ts">
	import { Greet } from '$wailsjs/go/backend/App';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card';
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';

	let name = $state('');
	let greeting = $state('');

	async function greet() {
		greeting = await Greet(name);
	}
</script>

<div class="flex min-h-screen items-center justify-center bg-background p-8">
	<div class="absolute top-4 right-4">
		<ThemeToggle />
	</div>
	<Card.Root class="w-full max-w-md">
		<Card.Header class="text-center">
			<Card.Title class="text-3xl font-bold tracking-tight">OmniBase</Card.Title>
			<Card.Description>Database Client</Card.Description>
		</Card.Header>
		<Card.Content class="space-y-4">
			<div class="flex gap-2">
				<Input
					type="text"
					bind:value={name}
					placeholder="Enter your name..."
				/>
				<Button onclick={greet}>Greet</Button>
			</div>

			{#if greeting}
				<p class="text-center text-lg font-medium text-foreground">{greeting}</p>
			{/if}
		</Card.Content>
	</Card.Root>
</div>
