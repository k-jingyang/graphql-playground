import gql from "graphql-tag";

export const getPokemon = gql`
  query pokemon($name: String!) {
    pokemon(name: $name) {
      id
      name
      image
    }
  }
`;
