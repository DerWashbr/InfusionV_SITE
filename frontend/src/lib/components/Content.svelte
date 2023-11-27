<script lang="ts">
	import Wave from '$components/Wave.svelte';
	import Hero from '$components/Hero.svelte';
	import Counter from '$components/Counter.svelte';
	import { fade } from 'svelte/transition';
	import { onMount } from 'svelte';

	let whyUsSection: HTMLElement | null;

	onMount(() => {
		whyUsSection = document.querySelector<HTMLElement>('#whyus');

		window.addEventListener('scroll', handleScroll);
	});

	function handleScroll() {
		const viewportHeight = window.innerHeight;
		const whyUsSectionOffset = whyUsSection?.getBoundingClientRect().top ?? 0;
		const scrollY = window.scrollY;

		if (whyUsSectionOffset < viewportHeight * 0.5 && scrollY < whyUsSectionOffset) {
			whyUsSection?.scrollIntoView({ behavior: 'smooth' });
			window.removeEventListener('scroll', handleScroll); // Entfernt den Event Listener nach dem Scrollen
		}
	}
</script>

<Counter />
<section id="welcome">
	<div class="container mx-auto p-8 space-y-8">
		<img class="w-180 h-180" src="https://i.imgur.com/ee4Note.png" alt="InfusionV Logo" />
	</div>

	<br />
	<br />

	<Wave inverted={false} />
</section>

<section id="whyus" transition:fade>
	<Hero />
</section>
