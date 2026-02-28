<script lang="ts">
	import SidePanelHeader from '$lib/components/sidebar/SidePanelHeader.svelte';
	import SidePanelToolbar from '$lib/components/sidebar/SidePanelToolbar.svelte';
	import ConnectionList from '$lib/components/sidebar/ConnectionList.svelte';
</script>

<div class="flex h-full w-60 flex-col border-r border-sidebar-border bg-sidebar">
	<SidePanelHeader />
	<SidePanelToolbar />
	<ConnectionList />
</div>
