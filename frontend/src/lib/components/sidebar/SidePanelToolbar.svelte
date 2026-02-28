<script lang="ts">
	import { Plus, RefreshCw, SlidersHorizontal, Search, FolderPlus, DatabaseZap, Upload } from 'lucide-svelte';
	import { connectionsStore } from '$lib/stores/connections';
	import { connectionDialog } from '$lib/stores/ui';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';

	const databaseTypes = [
		{ id: 'mysql', name: 'MySQL', icon: '🐬' },
		{ id: 'oracle', name: 'Oracle', icon: '🔴' },
		{ id: 'postgresql', name: 'PostgreSQL', icon: '🐘' },
		{ id: 'sqlserver', name: 'SQLServer', icon: '🔷' },
		{ id: 'mariadb', name: 'MariaDB', icon: '🔱' },
		{ id: 'clickhouse', name: 'ClickHouse', icon: '📊' },
		{ id: 'dm', name: 'DM', icon: '🔶' },
		{ id: 'presto', name: 'Presto', icon: '⚡' },
		{ id: 'db2', name: 'DB2', icon: '💠' },
		{ id: 'oceanbase', name: 'OceanBase', icon: '🌊' },
		{ id: 'oceanbase_oracle', name: 'OceanBase Oracle', icon: '🌊' },
		{ id: 'sqlite', name: 'SQLite', icon: '📦' },
		{ id: 'h2', name: 'H2', icon: '🟦' },
		{ id: 'hive', name: 'Hive', icon: '🐝' },
		{ id: 'kingbase', name: 'Kingbase', icon: '👑' }
	];

	function handleNewGroup() {
		connectionsStore.addGroup();
	}

	function handleSelectDbType(dbType: string) {
		connectionDialog.set({ open: true, driver: dbType, editId: null });
	}

	function handleImportConnections() {
		// TODO: implement import connections
	}
</script>

<div class="flex items-center gap-1 px-3 pb-2">
	<!-- Add button with dropdown -->
	<DropdownMenu.Root>
		<DropdownMenu.Trigger
			class="flex h-7 w-7 items-center justify-center rounded text-muted-foreground hover:bg-accent hover:text-foreground"
			title="Add"
		>
			<Plus size={16} />
		</DropdownMenu.Trigger>

		<DropdownMenu.Content class="w-52" align="start" sideOffset={4}>
			<!-- New group -->
			<DropdownMenu.Item class="gap-3" onclick={handleNewGroup}>
				<FolderPlus size={16} class="text-muted-foreground" />
				<span>New group</span>
			</DropdownMenu.Item>

			<!-- New connection submenu -->
			<DropdownMenu.Sub>
				<DropdownMenu.SubTrigger class="gap-3">
					<DatabaseZap size={16} class="text-muted-foreground" />
					<span>New connection</span>
				</DropdownMenu.SubTrigger>
				<DropdownMenu.SubContent class="max-h-[480px] w-52 overflow-y-auto">
					{#each databaseTypes as db}
						<DropdownMenu.Item class="gap-3" onclick={() => handleSelectDbType(db.id)}>
							<span class="w-5 text-center text-sm">{db.icon}</span>
							<span>{db.name}</span>
						</DropdownMenu.Item>
					{/each}
				</DropdownMenu.SubContent>
			</DropdownMenu.Sub>

			<!-- Import connections -->
			<DropdownMenu.Sub>
				<DropdownMenu.SubTrigger class="gap-3">
					<Upload size={16} class="text-muted-foreground" />
					<span>Import connections</span>
				</DropdownMenu.SubTrigger>
				<DropdownMenu.SubContent class="w-48">
					<DropdownMenu.Item onclick={handleImportConnections}>
						From file...
					</DropdownMenu.Item>
				</DropdownMenu.SubContent>
			</DropdownMenu.Sub>
		</DropdownMenu.Content>
	</DropdownMenu.Root>

	<button
		class="flex h-7 w-7 items-center justify-center rounded text-muted-foreground hover:bg-accent hover:text-foreground"
		title="Refresh"
	>
		<RefreshCw size={14} />
	</button>
	<button
		class="flex h-7 w-7 items-center justify-center rounded text-muted-foreground hover:bg-accent hover:text-foreground"
		title="Filter"
	>
		<SlidersHorizontal size={14} />
	</button>

	<!-- Search -->
	<button
		class="ml-auto flex h-7 flex-1 items-center gap-1.5 rounded-md border border-border bg-accent px-2 text-xs text-muted-foreground"
	>
		<Search size={12} />
		<span>Search</span>
		<kbd class="ml-auto rounded border border-border bg-accent px-1 py-0.5 text-[10px] text-muted-foreground/60">
			⌘F
		</kbd>
	</button>
</div>
