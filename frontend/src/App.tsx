import React from "react";

function App() {
  return (
    <div className="App">
      <input
        type="file"
        className="border-2 bg-slate-400"
        onChange={() => console.log("aaa")}
      />
    </div>
  );
}

export default App;
