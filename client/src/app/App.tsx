import React, { useState } from "react";

import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Pokemon from "../components/Pokemon/Pokemon";

const App = () => {
  const client = new ApolloClient({
    uri: "https://graphql-pokemon.now.sh",
  });

  return (
    <ApolloProvider client={client}>
      <div className="mx-auto h-full w-full flex justify-center flex-col items-center">
        <Aaaa />
      </div>
    </ApolloProvider>
  );
};

const Aaaa = () => {
  const [inputValue, setInputValue] = useState<string>("");
  const [pokemonName, setPokemonName] = useState<string>("pikachu");

  return (
    <React.Fragment>
      <form onSubmit={(e) => e.preventDefault()}>
        <div className="p-32">
          <label className="block text-gray-700">Enter a Pokemon Name</label>
          <input
            className="shadow appearance-none border rounded p-2 focus:outline-none"
            spellCheck={false}
            placeholder="e.g. pikachu"
            onChange={(e) => setInputValue(e.target.value)}
            value={inputValue}
          />
          <button
            className="bg-indigo-500 hover:bg-indigo-700 text-white py-2 px-4 m-2 rounded font-xs"
            onClick={() => setPokemonName(inputValue)}
          >
            Search
          </button>
        </div>

        <Pokemon pokemonName={pokemonName} />
      </form>
    </React.Fragment>
  );
};
export default App;
