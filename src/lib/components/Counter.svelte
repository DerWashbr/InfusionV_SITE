<script lang="ts">
	import { onMount } from 'svelte';

	let targetDate = new Date('2024-02-08').getTime();
	let currentDate = new Date().getTime();
	let remainingTime = targetDate - currentDate;

	let days = Math.floor(remainingTime / (1000 * 60 * 60 * 24));
	let hours = Math.floor((remainingTime % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
	let minutes = Math.floor((remainingTime % (1000 * 60 * 60)) / (1000 * 60));
	let seconds = Math.floor((remainingTime % (1000 * 60)) / 1000);

	let interval: NodeJS.Timeout;

	function startCountdown() {
		interval = setInterval(() => {
			if (days === 0 && hours === 0 && minutes === 0 && seconds === 0) {
				clearInterval(interval);
			} else {
				if (seconds > 0) {
					seconds--;
				} else {
					if (minutes > 0) {
						minutes--;
						seconds = 59;
					} else {
						if (hours > 0) {
							hours--;
							minutes = 59;
						} else {
							if (days > 0) {
								days--;
								hours = 23;
							}
						}
					}
				}
			}
		}, 1000);
	}

	onMount(() => {
		startCountdown();

		return () => {
			clearInterval(interval);
		};
	});
</script>

<div class="grid grid-flow-col bg-background justify-center gap-5 text-center auto-cols-max">
	<div class="flex flex-col p-2 bg-backgroundl rounded-box text-neutral-content">
		<span class="countdown font-mono text-5xl">
			<span style="--value:{days};"></span>
		</span>
		Tage
	</div>
	<div class="flex flex-col p-2 bg-background rounded-box text-neutral-content">
		<span class="countdown font-mono text-5xl">
			<span style="--value:{hours};"></span>
		</span>
		Stunden
	</div>
	<div class="flex flex-col p-2 bg-background rounded-box text-neutral-content">
		<span class="countdown font-mono text-5xl">
			<span style="--value:{minutes};"></span>
		</span>
		Minuten
	</div>
	<div class="flex flex-col p-2 bg-background rounded-box text-neutral-content">
		<span class="countdown font-mono text-5xl">
			<span style="--value:{seconds};"></span>
		</span>
		Sekunden
	</div>
</div>
