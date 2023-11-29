// src/auth/auth.ts
import OAuth from 'discord-oauth2';
import { PrismaClient } from '@prisma/client';

const clientId = process.env.DEINE_DISCORD_CLIENT_ID;
const clientSecret = process.env.DEIN_DISCORD_CLIENT_SECRET;
const redirectUri = process.env.DEINE_REDIRECT_URI;

// Erstellen einer Instanz für die Discord-OAuth2-Bibliothek
const oauth = new OAuth({
	clientId,
	clientSecret,
	redirectUri
});

// Erstellen einer Instanz von PrismaClient für die Datenbankoperationen
const prisma = new PrismaClient();

// Funktion zur Weiterleitung an Discord zur Authentifizierung
export const redirectToDiscord = () => {
	const authorizationUrl = oauth.generateAuthUrl({
		scope: ['identify', 'email'] // Discord-Berechtigungen, die du benötigst
	});
	return authorizationUrl;
};

// Funktion zur Überprüfung des OAuth-Callbacks und Rollen mit Prisma
export const handleDiscordCallback = async (code: string) => {
	try {
		// Austauschen des Authorization Codes gegen Access-Token und User-Informationen
		const tokenData = await oauth.tokenRequest({
			code,
			scope: ['identify', 'email'],
			grantType: 'authorization_code'
		});

		const discordUser = await oauth.getUser(tokenData.access_token);

		// Beispiel: Überprüfung der Rolle des Benutzers in der Datenbank mit Prisma
		const userRoles = await prisma.user.findUnique({
			where: { discordId: discordUser.id },
			select: { roles: true }
		});

		if (!userRoles || !userRoles.roles.includes('ROLE_X')) {
			// Wenn der Benutzer die erforderliche Rolle nicht hat, wird er ausgeloggt
			return null;
		}

		// Speichere oder aktualisiere die Benutzerdaten in der Datenbank nach Bedarf

		return discordUser; // Rückgabe der Discord-Benutzerdaten nach erfolgreicher Authentifizierung
	} catch (error) {
		console.error('Fehler bei der Verarbeitung des Discord-Callbacks:', error);
		return null;
	}
};

// Weitere Funktionen zum Speichern von Benutzerdaten oder Bereinigen von Sessions bei Bedarf
