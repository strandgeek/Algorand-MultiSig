import React, { FC } from 'react'

export interface HomeProps {}

export const Home: FC<HomeProps> = (props) => {
  return (
    <div className="flex items-center justify-center h-screen w-screen">
      Hello World
    </div>
  )
}
