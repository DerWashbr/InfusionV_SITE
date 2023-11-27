<script lang="ts">
	interface Streamer {
		Username: string;
	}

	import '@fontsource/coming-soon';
	import { onMount, afterUpdate } from 'svelte';
	// import '$css/twitch.pcss';
	import OfflineStream from '$components/twitch/OfflineStream.svelte';
	import OnlineStream from '$components/twitch/OnlineStream.svelte';
	import TwitchEmbed from '$components/twitch/TwitchEmbed.svelte';
	import {
		fetchData,
		isApiWorking,
		offlineStreamers,
		onlineStreamers,
		totalStreamers,
		totalViewers
	} from '$lib/api/streamers';
	let isMessage = false;
	let loadTime: Date | null = null;
	let currentStreamer: Streamer | null = null; // Definiere die Typdefinition für currentStreamer

	async function getData() {
		await fetchData();
		isApiWorking;
		offlineStreamers;
		onlineStreamers;
		totalStreamers;
		totalViewers;

		if (onlineStreamers.length > 0) {
			const randomIndex = Math.floor(Math.random() * onlineStreamers.length);
			currentStreamer = onlineStreamers[randomIndex];
		}
	}
	onMount(() => {
		getData();
		loadTime = new Date();
		// call getData every 4 minutes
		const interval = setInterval(
			() => {
				getData();
			},
			4 * 60 * 1000
		);
		// cleanup function to clear the interval when the component unmounts
		return () => clearInterval(interval);
	});
	afterUpdate(() => {
		if (loadTime) {
			const loadDuration = new Date().getTime() - loadTime.getTime();
			console.log(`Seite wurde in ${loadDuration}ms geladen.`);
		}
	});
</script>

<svelte:head>
	<title>InfusionV - More than Roleplay</title>
	<link rel="canonical" href="https://infusionV.de" />
</svelte:head>

<main>
	<div>
		{#await getData()}
			<div class="lds-heart"><div /></div>
		{:then}
			{#if !isApiWorking}
				<div role="alert" class="alert alert-error">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="stroke-current shrink-0 h-6 w-6"
						fill="none"
						viewBox="0 0 24 24"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
						/></svg
					>
					<span
						>Es gibt ein Problem mit der API. Bitte kontaktiere "DerWashbr" auf InfusionV - Danke</span
					>
				</div>
			{/if}
			<!-- {#each onlineStreamers as streamer}
				<div class="spotlight">
					<TwitchEmbed channel={streamer.Username} id="twitch-embed-{streamer.Username}" />
				</div>
			{/each} -->
			{#if onlineStreamers.length > 0 && offlineStreamers.length > 0}
				<div class="card shadow-lg">
					<div class="infobox">
						<center><div class="more-streams">Spotlight</div> </center>
					</div>
				</div>
				{#if currentStreamer}
					<div class="spotlight">
						<TwitchEmbed
							channel={currentStreamer.Username}
							id="twitch-embed-{currentStreamer.Username}"
						/>
					</div>
					<br />
				{/if}
			{/if}

			<!-- <img class="logo" src={logo} alt="Logo" /> -->
			{#if onlineStreamers.length > 0}
				<div class="card shadow-lg">
					<div class="infobox">
						<center
							><div class="more-streams">
								Es streamen gerade <b>{totalStreamers}</b> Streamer:innen und dort schauen
								<b>{totalViewers}</b> Zuschauer zu.
							</div>
						</center>
					</div>
				</div>
				<br />
			{/if}
			{#if isMessage}
				<div role="alert" class="alert alert-warning">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="stroke-current shrink-0 h-6 w-6"
						fill="none"
						viewBox="0 0 24 24"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
						/></svg
					>
					<span
						>Diese Seite ist in Bearbeitung! Bitte öffne ein Ticket auf InfusionV wenn du in die
						Liste eingepflegt werden willst - Danke</span
					>
				</div>
			{/if}
			{#if onlineStreamers.length > 0 && offlineStreamers.length > 0}
				<!-- <div class="totalplayer">Es sind gerade Spieler:innen auf InfusionV</div> -->
			{/if}
			<OnlineStream />
			<br />

			<div class="card shadow-lg">
				<div class="card-body">
					<center><div class="more-streams">Weitere Streamer:innen</div> </center>
				</div>
			</div>
			<br />
			<OfflineStream />
		{:catch error}
			<p>Error: {error.message}</p>
		{/await}
	</div>
</main>

<style>
	.more-streams {
		font-family:
			/* Coming Soon, */
			Lucida Grande,
			Lucida Sans Unicode,
			Lucida Sans,
			DejaVu Sans,
			Verdana,
			'sans-serif';
		font-weight: 400;
		line-height: 1.15em;
		margin: 0 0 15px;
		color: #4169e1;
		font-size: 26px;
	}
	.infobox {
		display: flex;
		flex: 1 1 auto;
		flex-direction: column;
		gap: -4.5rem;
	}
</style>
