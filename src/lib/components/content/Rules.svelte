<script lang="ts">
	import type { RegelnData } from '$lib/rules';
	import regelnData from '$lib/data/regeln.json';
	let regeln: RegelnData;
	$: {
		regeln = regelnData;
	}
</script>

<div id="regelbereiche">
	{#each regeln.regelbereiche as bereich (bereich.titel)}
		<div class="collapse collapse-plus bg-background">
			<input type="radio" name="my-accordion-3" />
			<div class="collapse-title text-xl font-medium"><center>{bereich.titel}</center></div>
			<div class="collapse-content">
				<ul>
					{#each bereich.regeln as regel (regel)}
						<li>{regel}</li>
					{/each}
				</ul>
			</div>
		</div>
	{/each}
</div>
