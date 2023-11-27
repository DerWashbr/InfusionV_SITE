<script>
	import SvelteSeo from 'svelte-seo';

	const PAGE_TITLE = 'InfusionV - More than Roleplay';
	const PAGE_URL = 'https://infusionV.de';
	const PAGE_DESC =
		'Entdecke das Beste an GTA V Roleplay auf InfusionV – dem einzigartigen Server seit 2024. Unvergessliche RP-Momente und eine tolle Community erwarten dich.';
</script>

<SvelteSeo
	title={PAGE_TITLE}
	description={PAGE_DESC}
	canonical={PAGE_URL}
	keywords="roleplay, gta rp server deutsch, "
	twitter={{
		image: 'https://infusionV.de/logo/infusionv.png',
		imageAlt: 'SnailyCAD',
		title: PAGE_TITLE,
		description: PAGE_DESC,
		card: 'summary'
	}}
	openGraph={{
		title: PAGE_TITLE,
		description: PAGE_DESC,
		url: PAGE_URL,
		type: 'website',
		images: [
			{
				url: 'https://infusionV.de/logo/infusionv.png',
				alt: 'InfusionV',
				width: 710,
				height: 710
			}
		]
	}}
/>
