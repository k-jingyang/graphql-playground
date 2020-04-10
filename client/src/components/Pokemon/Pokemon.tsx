import React from "react";
import { getPokemon } from "../../graphql/getPokemon";
import { motion } from "framer-motion";
import { useQuery } from "@apollo/react-hooks";

interface Pokemon {
  id: string;
  name: string;
  image: string;
}

interface PokemonData {
  pokemon: Pokemon;
}

interface Props {
  pokemonName: string;
}

const Pikachu = ({ pokemonName }: Props) => {
  const { loading, data } = useQuery<PokemonData>(getPokemon, { variables: { name: pokemonName } });

  return (
    <div className="text-gray-500 flex justify-center">
      {loading ? (
        <motion.div
          animate={{ y: 20 }}
          transition={{ yoyo: Infinity, duration: 0.5, type: "tween", ease: "easeInOut" }}
          key="loading"
        >
          Loading...{" "}
        </motion.div>
      ) : (
        <motion.div initial={{ opacity: 0 }} animate={{ opacity: 1 }} key="image">
          {data?.pokemon ? <img src={data?.pokemon?.image} alt="pikamon" /> : "Nothing found"}
        </motion.div>
      )}
    </div>
  );
};
export default Pikachu;
