// eslint-disable-next-line @typescript-eslint/no-var-requires
const { fontFamily } = require('tailwindcss/defaultTheme');
/** @type {import('tailwindcss').Config}*/
const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				background: '#1e2028',
				background2: '#121212',
				royalblue: '#4169e1'
			}
		},
		fontFamily: {
			sans: ['Raleway', ...fontFamily.sans],
			display: ['Raleway', ...fontFamily.sans],
			body: ['Raleway', ...fontFamily.sans]
		}
	},

	plugins: [require('daisyui')]
};

module.exports = config;
