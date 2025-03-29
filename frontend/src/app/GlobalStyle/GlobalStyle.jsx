import { createGlobalStyle } from 'styled-components';
import normalize from 'styled-normalize';

export const GlobalStyle = createGlobalStyle`
  ${normalize}
  html {
    font-size: 4px; //1rem
  }

  //названия для цветов тут: https://coolors.co/1a132d
  :root {
    --color-grape: #5829BB;
    --color-gray: #2e2e2e;
    --color-charcoal: #344054;
    --color-cadetGray: #98A2B3;
  }

  html, body, #root {
    width: 100%;
    height: 100%;
  }

  body {
    font-family: 'Inter', sans-serif;
    background: #151718;
  }
`;
