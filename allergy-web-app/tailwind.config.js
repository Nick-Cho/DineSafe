/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./src/components/*.{js,jsx,ts,tsx}"
  ], 
  theme: {
    colors: { 
        "white": "#ffffff",
        "black": "#000000",
        "btn-gray": "#333333",
        "gray": "#eeeeee"
    },
    fontFamily: {
      'uber': ['Uber Move Text', 'sans-serif']
    },
    extend: {
      gridTemplateColumns: {
        '16': 'repeat(16, minmax(0, 1fr))',        
      },
      gridColumnStart: {
        '13': '13',
        '14': '14',
        '15': '15',
        '16': '16',
        '17': '17',
      },
      margin: {
        '50%': '50%'
      },
      fontFamily: {
        'uber': ['Uber Move Text', 'sans-serif']
      }
    },
  },
  plugins: [],
}