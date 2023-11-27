<script lang="ts">
	import { onMount } from 'svelte';
	export let channel: string;
	export let width = '100%';
	export let height = '100%';
	export let autoplay = true;
    export let layout = 'video-with-chat';
	export let muted = true;
	export let parent = 'localhost';
	export let id: string;

	let username: Twitch.Player;

	onMount(async () => {
		username = new Twitch.Player(`twitch-embed-${id}`, {
			channel,
			width,
			height,
			autoplay,
            layout,
			muted,
			parent
		});
	});
</script>
<br />
<div id={`twitch-embed-${id}`} class="twitch-embed">
</div>

<style>
	.twitch-embed {
		width: 600px;
		aspect-ratio: 16/9;
		position: relative;
        margin: auto;
	}
</style>