<script>
	import '../app.pcss';
	import Navbar from '$components/Navbar.svelte';
	import Footer from '$components/Footer.svelte';
</script>

<main class="flex flex-col min-h-screen bg-background">
	<Navbar />
	<slot />
</main>
<Footer />

<style>
	:global(.main::-webkit-scrollbar) {
		width: 0.25rem;
	}

	:global(.main::-webkit-scrollbar-track) {
		background: #1e1e24;
	}

	:global(.main::-webkit-scrollbar-thumb) {
		color: #93caed;
	}

	:global(.main) {
		scrollbar-color: #93caed #1e1e24;
		scrollbar-width: 0.25rem;
	}
</style>
