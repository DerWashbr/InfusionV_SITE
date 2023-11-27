<script lang="ts">
	import { onlineStreamers, formatElapsedTime } from '$lib/api/streamers';
	import Icon from '@iconify/svelte';
	function goToTwitch(streamerName: string) {
		window.open(`https://twitch.tv/${streamerName}`, '_blank');
	}
</script>

<div class="flex justify-center">
	<div class="grid grid-cols-4 gap-6">
		{#each onlineStreamers as streamer}
			<div
				role="button"
				tabindex="0"
				on:click={() => goToTwitch(streamer.Username)}
				on:keydown={(event) => {
					if (event.key === 'Enter') {
						goToTwitch(streamer.Username);
					}
				}}
			>
				<div class="card w-96 bg-gray-800 shadow-xl max-h-[365px] max-w-[305px]">
					<img
						class="card-thumbnail"
						src={streamer.ThumbnailImg}
						alt="Stream Thumbnail"
						width="300"
						height="200"
						title={streamer.Title}
					/>
					<div class="card-body">
						<div class="avatar-online">
							<div class="w-16 mask mask-squircle">
								<img src={streamer.ProfileImg} alt={`${streamer.Username}'s Profile Image`} />
							</div>
						</div>
						<div class="badge-online badge badge-primary">LIVE</div>
						<h2 class="card-title">
							{streamer.Username}
						</h2>
						<div class="badge badge-fraktion badge-outline">{streamer.Fraktion}</div>
					</div>
					{#if streamer.LastStreamed}
						<div class="online-laststream-icon"><Icon icon="mdi-clock-time-four-outline" /></div>
						<div class="online-laststream-text">{formatElapsedTime(streamer.LastStreamed)}</div>
					{/if}
					<div class="online-viewer-icon"><Icon icon="mdi-eye" /></div>
					<div class="online-viewer-text">{streamer.Viewers}</div>
					<br />
					<!-- <div class="online-follower-icon"><Icon icon="mdi-account-plus" /></div>
						<div class="online-follower-text">{streamer.Followers}</div> -->
					<div class="online-charname">{streamer.Charname}</div>
					<div class="online-fraktion card-actions justify-center"></div>
				</div>
			</div>
		{/each}
	</div>
</div>
<br />

<style>
	.card-thumbnail {
		width: 100%;
		height: 180px;
		border-top-left-radius: 10px;
		border-top-right-radius: 10px;
	}

	div :global(.online-viewer-icon) {
		font-size: 32px;
		top: -160px;
		left: 24px;
		position: relative;
	}
	div :global(.online-follower-icon) {
		font-size: 32px;
		top: -154px;
		left: 206px;
		position: relative;
	}

	div :global(.online-laststream-icon) {
		font-size: 20px;
		position: relative;
		left: 211px;
		top: -109px;
	}
	.online-viewer-text {
		position: relative;
		left: 70px;
		top: -188px;
	}
	/* .online-follower-text {
		position: relative;
		left: 256px;
		top: -189px;
	} */
	.online-laststream-text {
		position: relative;
		left: 239px;
		top: -131px;
	}
	.online-charname {
		position: relative;
		left: 24px;
		top: -283px;
	}
	.online-fraktion {
		position: relative;
		left: 0px;
		top: -148px;
	}
	.avatar-online {
		position: relative;
		left: -11px;
		top: -64px;
	}
	.card-title {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		font-size: 1.25rem;
		line-height: 1.75rem;
		font-weight: 600;
		position: relative;
		left: -8px;
		top: -96px;
	}
	.badge-online {
		position: relative;
		left: 199px;
		top: -112px;
	}
	.badge-fraktion {
		position: relative;
		left: 22px;
		top: -13px;
	}
</style>
