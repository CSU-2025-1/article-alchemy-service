import {StyleSheetManager} from "@/app/StyleSheetManager";
import {GlobalStyle} from "@/app/GlobalStyle";
import {Router} from "@/app/Router";


function App() {
  return <StyleSheetManager>
    <GlobalStyle />
    <Router />
  </StyleSheetManager>;
}

export default App;
