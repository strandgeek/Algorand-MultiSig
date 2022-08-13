import React, { FC } from "react";
import { useNavigate } from "react-router-dom";
import { useMeQuery } from "../../client/queries";
import { getIdenticonSrc } from "../../utils/getIdenticonSrc";
import { getShortAddress } from "../../utils/getShortAddress";

export interface NavbarProps {}

export const Navbar: FC<NavbarProps> = (props) => {
  const navigate = useNavigate()
  const { data: me } = useMeQuery()
  const meAvatar = getIdenticonSrc(me?.address)
  const logout = () => {
    localStorage.removeItem('token')
    navigate('/')
  }
  return (
    <div className="bg-base-100 border-b border-b-base-300">
      <div className="navbar max-w-8xl mx-auto">
        <div className="navbar-start">
          <div className="flex-1">
            <a className="btn btn-ghost normal-case text-xl">TxSigner</a>
          </div>
        </div>
        <div className="navbar-center hidden lg:flex">
          <ul className="menu menu-horizontal p-0">
            <li>
              <a>Item 1</a>
            </li>
            <li>
              <a>Item 2</a>
            </li>
            <li>
              <a>Item 3</a>
            </li>
          </ul>
        </div>
        <div className="navbar-end">
          <div className="flex-none">
            <div className="dropdown dropdown-end">
              <label tabIndex={0} className="btn btn-ghost avatar flex items-center">
                <div className="w-10 rounded-full mr-2">
                  <img src={meAvatar} />
                </div>
                {getShortAddress(me?.address)}
              </label>
              <ul
                tabIndex={0}
                className="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
              >
                <li>
                  <button onClick={() => logout()}>
                    Logout
                  </button>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
