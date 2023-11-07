class grafo:
    def __init__(self, dirigido=True, vertices=None):
        self.dirigido = dirigido
        self.datos = {}
        if vertices is not None:
            for i in vertices: self.datos[i] = {}

    def agregar_arista(self, origen, destino, peso=1):
        dicDestino = self.datos.get(destino, {})
        dicOrigen = self.datos.get(origen, {})
        dicOrigen[destino] = peso
        self.datos[origen] = dicOrigen
        if not self.dirigido:
            dicDestino[origen] = peso
        self.datos[destino] = dicDestino

    def agregar_vertice(self, dato):
        dic = self.datos.get(dato, {})
        self.datos[dato] = dic

    def adyacentes(self, vertice):
        return self.datos[vertice].keys()

    def obtener_vertices(self):
        return list(self.datos.keys())

    def obtener_aristas(self):
        vertices = []
        visitados = set()
        for origen, destinos in self.datos.items():
            for destino, peso in destinos.items():
                if (origen, destino) not in visitados and (destino, origen) not in visitados:
                    vertices.append((origen, destino, peso))
                    visitados.add((origen, destino))
        return vertices

    def vertice_aleatorio(self):
        for elem in self.datos.keys():
            return elem

    def peso(self, origen, destino):
        dic = self.datos[origen]
        return dic[destino]

    def pertenece(self, dato):
        if dato in self.datos: return True
        return False

    def unidos(self, origen, destino):
        if origen not in self.datos or destino not in self.datos: return False
        if destino in self.datos[origen] or origen in self.datos[destino]: return True
        return False

    def __iter__(self):
        self.iterador = list(self.datos.keys())
        self.actual = 0
        return self

    def __next__(self):
        if self.actual >= len(self.iterador): raise StopIteration
        dato = self.iterador[self.actual]
        self.actual += 1
        return dato

    def __len__(self):
        return len(self.datos.keys())
