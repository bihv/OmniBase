<script lang="ts">
	import {
		Database,
		SquareTerminal,
		ChartColumn,
		Bot,
		RefreshCw,
		Bell,
		MessageSquare,
		Shield
	} from 'lucide-svelte';

	let activeItem = $state('console');

	const topItems = [
		{ id: 'database', icon: Database, label: 'Database' },
		{ id: 'console', icon: SquareTerminal, label: 'SQL Console' },
		{ id: 'charts', icon: ChartColumn, label: 'Charts' },
		{ id: 'ai', icon: Bot, label: 'AI Assistant' }
	];

	const bottomItems = [
		{ id: 'refresh', icon: RefreshCw, label: 'Refresh' },
		{ id: 'notifications', icon: Bell, label: 'Notifications' },
		{ id: 'chat', icon: MessageSquare, label: 'Chat' },
		{ id: 'settings', icon: Shield, label: 'Settings' }
	];
</script>

<div class="flex h-full w-12 flex-col items-center bg-card border-r border-sidebar-border py-3">
	<!-- Logo -->
	<button
		class="mb-4 flex h-8 w-8 items-center justify-center rounded-lg bg-primary text-sm font-bold text-primary-foreground"
	>
		O
	</button>

	<!-- Top nav items -->
	<div class="flex flex-1 flex-col items-center gap-1">
		{#each topItems as item}
			<button
				class="flex h-10 w-10 items-center justify-center rounded-lg transition-colors {activeItem === item.id
					? 'bg-primary/20 text-primary'
					: 'text-muted-foreground hover:bg-accent hover:text-foreground'}"
				title={item.label}
				onclick={() => (activeItem = item.id)}
			>
				<item.icon size={20} />
			</button>
		{/each}
	</div>

	<!-- Bottom nav items -->
	<div class="flex flex-col items-center gap-1">
		{#each bottomItems as item}
			<button
				class="flex h-10 w-10 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:bg-accent hover:text-foreground"
				title={item.label}
			>
				<item.icon size={20} />
			</button>
		{/each}
	</div>
</div>
