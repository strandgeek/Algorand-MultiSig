import React, { FC } from 'react'
import { Navbar } from './Navbar'

export interface AppLayoutProps {
  children: React.ReactNode
}

export const AppLayout: FC<AppLayoutProps> = ({ children }) => {
  return (
    <div className="bg-base-200 h-screen">
      <Navbar />
      {children}
    </div>
  )
}
