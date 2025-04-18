
import TodoList from "@/components/todos/TodoList";

export default function Index() {
  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-4xl mx-auto">
        <header className="text-center mb-8">
          <h1 className="text-3xl font-bold text-gray-900">My Todos</h1>
          <p className="text-gray-600 mt-2">Organize your tasks efficiently</p>
        </header>
        <TodoList />
      </div>
    </div>
  );
}
