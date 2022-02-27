import React, { createContext, useContext, useState } from "react";
const windowTS = window as any

export interface ConfigContextType {
  title: string
  blockExplorerName: string
  headBlockNumberApp: string
}

export const ConfigContext = createContext<ConfigContextType>(null!)


export default function ConfigProvider(props: React.PropsWithChildren<{}>) {
  const title = (windowTS.TopLevelConfig ? windowTS.TopLevelConfig.title  : "中数文联盟链历史节点管理台")
  const blockExplorerName = (windowTS.TopLevelConfig ? windowTS.TopLevelConfig.blockExplorerName  : "区块浏览器")
  const headBlockNumberApp = (windowTS.TopLevelConfig ? windowTS.TopLevelConfig.headBlockNumberApp  : "历史节点系统")

  return (
    <ConfigContext.Provider value={{title, blockExplorerName, headBlockNumberApp}}>
      {props.children}
    </ConfigContext.Provider>
  )
}

export function useConfig() {
  const context = useContext(ConfigContext)
  if(!context) {
    throw new Error("usConfig must be used within the ConfigProvider")
  }
  return context
}