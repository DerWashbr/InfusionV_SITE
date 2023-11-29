<script lang="ts">
	import Icon from '@iconify/svelte';
	let username = '';
	let password = '';
	let email = '';
	let loading = false;
	let success = false;
	let error: string | null = null;
	let showForgotPassword = false;
	let showSignIn = false;

	async function handleSubmit() {
		try {
			loading = true;

			const response = await fetch('http://localhost:7080/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username,
					password
				})
			});
			const data = await response.json();
			if (response.ok) {
				console.log('JWT-Token:', data.token);
				success = true;
			} else {
				throw new Error(data.message || 'Unbekannter Fehler');
			}
		} catch (error: any) {
			console.error(error);
			error = error.message || 'Unbekannter Fehler';
		} finally {
			loading = false;
		}
	}

	async function handleRegister() {
		try {
			loading = true;

			const response = await fetch('http://localhost:7080/register', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username,
					email,
					password
				})
			});

			const data = await response.json();
			if (response.ok) {
				console.log('Registrierung erfolgreich:', data.message);
				success = true;
			} else {
				throw new Error(data.message || 'Unbekannter Fehler');
			}
		} catch (error: any) {
			console.error(error);
			error = error.message || 'Unbekannter Fehler';
		} finally {
			loading = false;
		}
	}

	async function handleChangePassword() {
		try {
			loading = true;

			const response = await fetch('http://localhost:7080/changepassword', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username,
					password
				})
			});

			const data = await response.json();
			if (response.ok) {
				console.log('Passwort erfolgreich geändert:', data.message);
				success = true;
			} else {
				throw new Error(data.message || 'Unbekannter Fehler');
			}
		} catch (error: any) {
			console.error(error);
			error = error.message || 'Unbekannter Fehler';
		} finally {
			loading = false;
		}
	}

	$: {
		if (success) {
			setTimeout(() => {
				success = false;
			}, 3000);
		}
	}
</script>

<div class="flex justify-center items-center h-screen">
	<div class="card w-96 bg-background2 shadow-xl">
		<div class="card-body">
			<h2 class="card-title justify-center text-royalblue">Admincontrolpanel</h2>
<br />
			{#if !showForgotPassword && !showSignIn}
				<form on:submit|preventDefault={handleSubmit} class="flex flex-col items-center">
					<input
						type="username"
						placeholder="Benutzername"
						class="input input-bordered input-info w-full max-w-xs"
						bind:value={username}
					/>
					<br />
					<input
						type="password"
						placeholder="Passwort"
						class="input input-bordered input-info w-full max-w-xs"
						bind:value={password}
					/>
					<br />
					<div class="flex justify-between w-full">
						<input type="submit" value="Einloggen" class="btn btn-background" />
						<button type="submit" class="btn btn-disabled">
							<Icon icon="akar-icons:discord-fill" width="24" height="24" />
						</button>
					</div>
					{#if loading}
						<p>Wird geladen...</p>
					{:else if success}
						<p>Erfolgreich eingeloggt!</p>
					{:else if error}
						<p class="text-red-500">Fehler: {error}</p>
					{/if}
				</form>

				<button on:click={() => (showForgotPassword = true)} class="btn mt-4">
					Passwort vergessen?
				</button>

				<button on:click={() => (showSignIn = true)} class="btn mt-4">
					Registrieren
				</button>
			{/if}

			{#if showForgotPassword}
				<form
					on:submit|preventDefault={handleChangePassword}
					class="flex flex-col items-center"
				>
					<input
						type="email"
						placeholder="E-Mail-Adresse"
						class="input input-bordered input-info w-full max-w-xs"
						bind:value={username}
					/>

					<input
						type="submit"
						value="Passwort zurücksetzen"
						class="btn mt-4 btn-secondary"
					/>

					{#if loading}
						<p>Wird geladen...</p>
					{:else if success}
						<p>Passwort wurde zurückgesetzt!</p>
					{:else if error}
						<p class="text-red-500">Fehler: {error}</p>
					{/if}
				</form>

				<button on:click={() => (showForgotPassword = false)} class="btn mt-4">
					Zurück zum Login
				</button>
			{/if}

			{#if showSignIn}
				<form on:submit|preventDefault={handleRegister} class="flex flex-col items-center">
					<input
						type="username"
						placeholder="Benutzername"
						class="input input-bordered input-info w-full max-w-xs"
						bind:value={username}
					/>
					<br />
					<input
						type="password"
						placeholder="Passwort"
						class="input input-bordered input-info w-full max-w-xs"
						bind:value={password}
					/> <br />
					<input
						type="E-Mail"
						placeholder="E-Mail-Adresse"
						class="input input-bordered input-info w-full max-w-xs"
						bind:value={email}
					/>

					<input type="submit" value="Registrieren" class="btn mt-4" />

					{#if loading}
						<p>Wird geladen...</p>
					{:else if success}
						<p>Erfolgreich angemeldet!</p>
					{:else if error}
						<p class="text-red-500">Fehler: {error}</p>
					{/if}
				</form>

				<button on:click={() => (showSignIn = false)} class="btn mt-4">
					Zurück zum Login
				</button>
			{/if}
		</div>
	</div>
</div>

<style>
	.input-info {
		border: 2px solid #4169e1;
	}
/* .btn-primary {
    --tw-text-opacity: 1;
  background-color: #4169e1;
  color: #121212;
} */

.btn-background {
	color: #121212;
    background-color: royalblue;
}
</style>
