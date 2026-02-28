<script lang="ts">
	import ActivityBar from '$lib/components/layout/ActivityBar.svelte';
	import SidePanel from '$lib/components/layout/SidePanel.svelte';
	import MainContent from '$lib/components/layout/MainContent.svelte';
	import RightBar from '$lib/components/layout/RightBar.svelte';
	import StatusBar from '$lib/components/layout/StatusBar.svelte';
	import AddConnectionForm from '$lib/components/connection/AddConnectionForm.svelte';
</script>

<div class="flex h-screen w-screen flex-col overflow-hidden bg-background">
	<!-- Main area -->
	<div class="flex flex-1 overflow-hidden">
		<!-- Activity Bar (leftmost icon strip) -->
		<ActivityBar />

		<!-- Side Panel (database tree) -->
		<SidePanel />

		<!-- Main Content Area -->
		<MainContent />

		<!-- Right Bar -->
		<RightBar />
	</div>

	<!-- Status Bar -->
	<StatusBar />
</div>

<!-- Connection Dialog (always mounted, controlled by uiStore) -->
<AddConnectionForm />
