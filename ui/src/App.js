import "./App.css";
import { store } from "./app/store";
import { Provider } from "react-redux";
import { DashboardLayout } from "./components/dashboard-layout";
import { ThemeProvider } from "@mui/material/styles";
import { theme } from "./theme";

function App() {
  return (
    <div className="App">
      <Provider store={store}>
        <ThemeProvider theme={theme}>
          <DashboardLayout />
        </ThemeProvider>
      </Provider>
    </div>
  );
}

export default App;
