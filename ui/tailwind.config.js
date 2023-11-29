/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
        fontSize: {
            mono: "'JetBrains Mono', monospace"
        }
    },
  },
  plugins: [],
}

