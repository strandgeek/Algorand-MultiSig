import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import {
  QueryClientProvider,
} from '@tanstack/react-query'
import { Home } from "./pages/Home";
import { queryClient } from "./client";

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
        </Routes>
      </Router>
    </QueryClientProvider>
  );
}

export default App;
