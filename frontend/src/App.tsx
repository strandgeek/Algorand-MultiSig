import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import {
  QueryClientProvider,
} from '@tanstack/react-query'
import { queryClient } from "./client";

// Pages
import { ListMultiSigAccounts } from "./pages/app/ListMultiSigAccounts";
import { Index } from "./pages/Index";

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Router>
        <Routes>
          <Route path="/" element={<Index />} />
          <Route path="/app" element={<ListMultiSigAccounts />} />
        </Routes>
      </Router>
    </QueryClientProvider>
  );
}

export default App;
