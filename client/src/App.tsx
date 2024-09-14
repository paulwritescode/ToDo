import { Center, Container, Image, Stack, Text } from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

export const BASE_URL =
  import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";
function App() {
  return (
    <>
      <Stack h="80vh">
        <Navbar />
        <Container>
          <TodoForm />
          <TodoList />
        </Container>
      </Stack>
      <Center>
        <Stack direction="row">
          <Image
            borderRadius="full"
            boxSize="100px"
            src="/todo.png"
            alt="Dan Abramov"
          />
          <Text fontSize="sm">
            <a target="_blank" href="https://paulmbugua.netlify.app/">
              2024 Paul (moon) <br />
              To check out my work
            </a>
          </Text>
          <a target="_blank" href="https://paulmbugua.netlify.app/">
            <Image
              borderRadius="full"
              boxSize="100px"
              src="/tajiri-final.png"
              alt="Dan Abramov"
            />
          </a>
        </Stack>
      </Center>
    </>
  );
}

export default App;
