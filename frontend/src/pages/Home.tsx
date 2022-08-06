import React, { FC } from 'react'
import { AppLayout } from '../layouts/AppLayout'

export interface HomeProps {}

export const Home: FC<HomeProps> = (props) => {
  return (
    <AppLayout>
      <div className="flex items-center justify-center h-screen w-screen">
        Hello World
      </div>
    </AppLayout>
  )
}
