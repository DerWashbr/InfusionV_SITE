<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { afterUpdate } from 'svelte';

	let isSticky = false;
	let firstRender = true;

	const handleScroll = () => {
		const currentScrollY = window.scrollY;
		isSticky = currentScrollY > 0;
	};
	const optimizedScrollHandler = () => {
		let ticking = false;
		return () => {
			if (!ticking) {
				window.requestAnimationFrame(() => {
					handleScroll();
					ticking = false;
				});
				ticking = true;
			}
		};
	};

	onMount(() => {
		window.addEventListener('scroll', optimizedScrollHandler());

		return () => {
			window.removeEventListener('scroll', optimizedScrollHandler());
		};
	});
	afterUpdate(() => {
		// Überprüfen, ob es sich um das erste Rerendering handelt
		if (firstRender) {
			firstRender = false;
			isSticky = window.scrollY > 0;
		}
	});
</script>

{#if !firstRender}
	<div class:sticky={isSticky} class="navbar mb-2 shadow-lg text-royalblue bg-background2">
		<div class="px-2 mx-2 navbar-start">
			<span class="text-lg font-bold"> InfusionV </span>
		</div>
		<div class="hidden px-2 mx-2 navbar-center lg:flex">
			<div class="flex items-stretch">
				<a href="/" class="btn btn-ghost btn-sm rounded-btn"> Home </a>
				<a href="rules" class="btn btn-ghost btn-sm rounded-btn"> Regeln </a>
				<a href="discord" class="btn btn-ghost btn-sm rounded-btn"> Discord </a>
				<a href="twitch" class="btn btn-ghost btn-sm rounded-btn"> Jetzt Live! </a>
			</div>
		</div>
		<div class="navbar-end">
			<p class="btn btn-ghost btn-sm rounded-btn">Spieler: 0</p>
			<a href="/user">
				<button class="btn">
					<Icon icon="akar-icons:person" width="24" height="24" />
				</button>
			</a>
		</div>
	</div>
{/if}

<style>
	.sticky {
		position: fixed;
		top: 0;
		width: 100%;
		z-index: 1000;
	}
</style>
