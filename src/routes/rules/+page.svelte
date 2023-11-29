<script>
	import Seo from '$components/Seo.svelte';
	import Rules from '$components/content/Rules.svelte';
</script>

<Seo />

<svelte:head>
	<title>InfusionV - More than Roleplay</title>
	<link rel="canonical" href="https://infusionV.de" />
</svelte:head>

<div class="card shadow-lg">
	<div class="card-body">
		<center><div class="rules-desc">Unsere Regeln</div> </center>
	</div>
</div>
<Rules />

<style>
	.rules-desc {
			/* Coming Soon, */
    font-family: Lucida Grande, Lucida Sans Unicode, Lucida Sans, DejaVu Sans, Verdana, 'sans-serif';
    font-weight: 400;
    line-height: 1.15em;
    margin: 0 0 15px;
    color: #4169e1;
    font-size: 26px;
    position: relative;
    left: -12px;
	}
</style>
