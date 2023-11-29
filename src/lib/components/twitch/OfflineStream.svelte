<script lang="ts">
	import { offlineStreamers, formatElapsedTime } from '$lib/api/streamers';
</script>

<div class="flex justify-center">
	<div class="grid grid-cols-5 gap-6">
		{#each offlineStreamers as streamer}
			<div class="card w-96 bg-base-100 shadow-xl max-w-[320px]">
				<div class="card-body">
					<div class="avatar">
						<div class="w-16 mask mask-squircle">
							<img src={streamer.ProfileImg} alt={`${streamer.Username}'s Profile Image`} />
						</div>
						<h2 class="card-title mx-auto">{streamer.Username}</h2>
					</div>
					<div class="card-charname">{streamer.Charname}</div>
					<br />
					{#if streamer.LastStreamed}
						<div class="card-time-text-off">
							Letzter Stream: {formatElapsedTime(streamer.LastStreamed)}
						</div>
					{/if}
					<div class="card-actions justify-center">
						<div class="badge badge-outline">{streamer.Fraktion}</div>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>

<style>
	.card-charname {
		display: flex;
		justify-content: center;
		top: 28px;
		position: relative;
	}
</style>
