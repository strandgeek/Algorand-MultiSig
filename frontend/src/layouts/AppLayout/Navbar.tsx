import { useQueryClient } from "@tanstack/react-query";
import React, { FC } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useMeQuery } from "../../client/queries";
import { getIdenticonSrc } from "../../utils/getIdenticonSrc";
import { getShortAddress } from "../../utils/getShortAddress";

export interface NavbarProps {}

export const Navbar: FC<NavbarProps> = (props) => {
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const { data: me } = useMeQuery()
  const meAvatar = getIdenticonSrc(me?.address)
  const logout = () => {
    localStorage.removeItem('token')
    queryClient.clear()
    navigate('/')
  }
  return (
    <div className="bg-base-100 border-b border-b-base-300">
      <div className="navbar max-w-8xl mx-auto">
        <div className="navbar-start">
          <div className="flex-1">
            <Link to="/app" className="btn btn-ghost normal-case text-xl">
              <img src="/logo.png" className="h-6 w-6 mr-1" alt="TxSigner" />
              TxSigner
            </Link>
          </div>
        </div>
        <div className="navbar-end">
          <div className="flex-none">
            <div className="dropdown dropdown-end">
              <label tabIndex={0} className="btn btn-ghost avatar flex items-center">
                <div className="w-10 rounded-full mr-2">
                  <img src={meAvatar} alt="User Avatar" />
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
