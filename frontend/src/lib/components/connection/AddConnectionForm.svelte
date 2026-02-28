<script lang="ts">
	import {
		TestTube,
		Save,
		Loader2,
		Eye,
		EyeOff,
		CheckCircle,
		XCircle
	} from 'lucide-svelte';
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';
	import { Label } from '$lib/components/ui/label';
	import { Separator } from '$lib/components/ui/separator';
	import * as Select from '$lib/components/ui/select';
	import * as Alert from '$lib/components/ui/alert';
	import * as Dialog from '$lib/components/ui/dialog';
	import { connectionDialog } from '$lib/stores/ui';
	import { TestConnection, SaveConnection } from '../../../../wailsjs/go/controllers/ConnectionController';

	const driverLabels: Record<string, string> = {
		mysql: 'MySQL',
		oracle: 'Oracle',
		postgresql: 'PostgreSQL',
		sqlserver: 'SQL Server',
		mariadb: 'MariaDB',
		clickhouse: 'ClickHouse',
		dm: 'DM',
		presto: 'Presto',
		db2: 'DB2',
		oceanbase: 'OceanBase',
		oceanbase_oracle: 'OceanBase Oracle',
		sqlite: 'SQLite',
		h2: 'H2',
		hive: 'Hive',
		kingbase: 'Kingbase'
	};

	const driverIcons: Record<string, string> = {
		mysql: '🐬',
		oracle: '🔴',
		postgresql: '🐘',
		sqlserver: '🔷',
		mariadb: '🔱',
		clickhouse: '📊',
		dm: '🔶',
		presto: '⚡',
		db2: '💠',
		oceanbase: '🌊',
		oceanbase_oracle: '🌊',
		sqlite: '📦',
		h2: '🟦',
		hive: '🐝',
		kingbase: '👑'
	};

	const defaultPorts: Record<string, number> = {
		mysql: 3306,
		oracle: 1521,
		postgresql: 5432,
		sqlserver: 1433,
		mariadb: 3306,
		clickhouse: 8123,
		dm: 5236,
		presto: 8080,
		db2: 50000,
		oceanbase: 2881,
		oceanbase_oracle: 2881,
		sqlite: 0,
		h2: 9092,
		hive: 10000,
		kingbase: 54321
	};

	const sslOptions = [
		{ value: 'disable', label: 'Disable' },
		{ value: 'require', label: 'Require' },
		{ value: 'verify-ca', label: 'Verify CA' },
		{ value: 'verify-full', label: 'Verify Full' },
		{ value: 'prefer', label: 'Prefer' }
	];

	// Reactive properties from the store
	let driver = $derived($connectionDialog.driver || 'mysql');
	let editId = $derived($connectionDialog.editId);
	let dialogOpen = $derived($connectionDialog.open);
	let isSQLite = $derived(driver === 'sqlite');

	// Form state
	let name = $state('');
	let host = $state('localhost');
	let port = $state(3306);
	let user = $state('root');
	let password = $state('');
	let database = $state('');
	let sslMode = $state<string | undefined>('disable');
	let showPassword = $state(false);

	// Status state
	let testing = $state(false);
	let saving = $state(false);
	let testResult = $state<{ success: boolean; message: string } | null>(null);

	// Reset form whenever dialog opens with a new driver
	$effect(() => {
		if (dialogOpen) {
			name = '';
			host = 'localhost';
			port = defaultPorts[driver] ?? 3306;
			user = 'root';
			password = '';
			database = '';
			sslMode = 'disable';
			showPassword = false;
			testResult = null;
		}
	});

	function handleOpenChange(open: boolean) {
		if (!open) connectionDialog.set({ open: false, driver: '', editId: null });
	}

	function buildConfig() {
		return {
			driver,
			host,
			port,
			user,
			password,
			database,
			sslMode: sslMode ?? 'disable'
		};
	}

	async function handleTestConnection() {
		testing = true;
		testResult = null;
		try {
			const result = await TestConnection(buildConfig());
			testResult = { success: true, message: result };
		} catch (err: any) {
			testResult = { success: false, message: err?.message || String(err) };
		} finally {
			testing = false;
		}
	}

	async function handleSave() {
		saving = true;
		try {
			const connInfo = {
				id: editId || '',
				name: name || `${driverLabels[driver] || driver} - ${host}`,
				driver,
				host,
				port,
				user,
				password,
				database,
				sslMode: sslMode ?? 'disable'
			};
			await SaveConnection(connInfo);
			connectionDialog.set({ open: false, driver: '', editId: null });
		} catch (err: any) {
			testResult = { success: false, message: `Save failed: ${err?.message || String(err)}` };
		} finally {
			saving = false;
		}
	}
</script>

<Dialog.Root open={dialogOpen} onOpenChange={handleOpenChange}>
	<Dialog.Content class="sm:max-w-2xl max-h-[85vh] flex flex-col gap-0 p-0">
		<!-- Header -->
		<Dialog.Header class="px-6 pt-6 pb-4">
			<Dialog.Title class="flex items-center gap-2">
				<span class="text-lg">{driverIcons[driver] || '🔌'}</span>
				{editId ? 'Edit' : 'New'} {driverLabels[driver] || driver} Connection
			</Dialog.Title>
			<Dialog.Description>
				Configure your {driverLabels[driver] || driver} database connection details
			</Dialog.Description>
		</Dialog.Header>

		<!-- Scrollable form body -->
		<div class="flex-1 overflow-y-auto px-6 pb-2">
			<div class="grid gap-6">
				<!-- Connection Name -->
				<div class="grid gap-2">
					<Label for="conn-name">Connection Name</Label>
					<Input
						id="conn-name"
						bind:value={name}
						placeholder={`My ${driverLabels[driver] || driver} Connection`}
					/>
					<p class="text-xs text-muted-foreground">
						A friendly name to identify this connection
					</p>
				</div>

				{#if !isSQLite}
					<Separator />

					<!-- Host & Port -->
					<div class="grid grid-cols-3 gap-4">
						<div class="col-span-2 grid gap-2">
							<Label for="conn-host">Host</Label>
							<Input id="conn-host" bind:value={host} placeholder="localhost" />
						</div>
						<div class="grid gap-2">
							<Label for="conn-port">Port</Label>
							<Input
								id="conn-port"
								type="number"
								bind:value={port}
								placeholder={String(defaultPorts[driver] ?? '')}
							/>
						</div>
					</div>

					<Separator />

					<!-- Authentication -->
					<div class="grid grid-cols-2 gap-4">
						<div class="grid gap-2">
							<Label for="conn-user">Username</Label>
							<Input id="conn-user" bind:value={user} placeholder="root" />
						</div>
						<div class="grid gap-2">
							<Label for="conn-password">Password</Label>
							<div class="relative">
								<Input
									id="conn-password"
									type={showPassword ? 'text' : 'password'}
									bind:value={password}
									placeholder="••••••••"
									class="pr-10"
								/>
								<Button
									variant="ghost"
									size="icon-sm"
									class="absolute right-1 top-1/2 h-7 w-7 -translate-y-1/2"
									onclick={() => (showPassword = !showPassword)}
								>
									{#if showPassword}
										<EyeOff size={16} />
									{:else}
										<Eye size={16} />
									{/if}
								</Button>
							</div>
						</div>
					</div>
				{/if}

				<Separator />

				<!-- Database -->
				<div class="grid gap-2">
					<Label for="conn-database">{isSQLite ? 'Database File Path' : 'Database'}</Label>
					<Input
						id="conn-database"
						bind:value={database}
						placeholder={isSQLite ? '/path/to/database.db' : 'mydb'}
					/>
				</div>

				{#if !isSQLite}
					<Separator />

					<!-- SSL Mode -->
					<div class="grid gap-2">
						<Label>SSL Mode</Label>
						<Select.Root type="single" bind:value={sslMode}>
							<Select.Trigger class="w-full">
								{sslOptions.find((o) => o.value === sslMode)?.label ?? 'Select SSL mode'}
							</Select.Trigger>
							<Select.Content>
								{#each sslOptions as opt}
									<Select.Item value={opt.value} label={opt.label} />
								{/each}
							</Select.Content>
						</Select.Root>
					</div>
				{/if}

				<!-- Test result -->
				{#if testResult}
					{#if testResult.success}
						<Alert.Root>
							<CheckCircle size={16} />
							<Alert.Title>Connection Successful</Alert.Title>
							<Alert.Description>{testResult.message}</Alert.Description>
						</Alert.Root>
					{:else}
						<Alert.Root variant="destructive">
							<XCircle size={16} />
							<Alert.Title>Connection Failed</Alert.Title>
							<Alert.Description>{testResult.message}</Alert.Description>
						</Alert.Root>
					{/if}
				{/if}
			</div>
		</div>

		<!-- Footer -->
		<Dialog.Footer class="px-6 py-4 border-t border-border">
			<Button variant="outline" onclick={handleTestConnection} disabled={testing}>
				{#if testing}
					<Loader2 size={16} class="animate-spin" />
				{:else}
					<TestTube size={16} />
				{/if}
				Test Connection
			</Button>
			<Button onclick={handleSave} disabled={saving}>
				{#if saving}
					<Loader2 size={16} class="animate-spin" />
				{:else}
					<Save size={16} />
				{/if}
				Save
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
