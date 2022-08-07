import React from "react";
import { ToastContainer } from 'react-toastify';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import {
  QueryClientProvider,
} from '@tanstack/react-query'
import { queryClient } from "./client";

// Pages
import { ListMultiSigAccounts } from "./pages/app/ListMultiSigAccounts";
import { CreateMultiSigAccount } from "./pages/app/CreateMultiSigAccount";
import { Index } from "./pages/Index";

function App() {
  return (
    <>
      <ToastContainer />
      <QueryClientProvider client={queryClient}>
        <Router>
          <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/app" element={<ListMultiSigAccounts />} />
            <Route path="/app/create-multisig-account" element={<CreateMultiSigAccount />} />
          </Routes>
        </Router>
      </QueryClientProvider>
    </>
  );
}

export default App;
