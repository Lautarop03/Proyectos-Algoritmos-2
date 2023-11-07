import csv
import heapq
from collections import deque
from grafo import grafo

ESCALA = "camino_escalas"
TOTAL = "TOTAL"


# Funciones de carga de archivos
def cargarAeropuertos(ruta, g, ciudades, coordenadas):
    with open(ruta) as f:
        for linea in f:
            campos = linea.split(",")
            ciudad, codigo_aeropuerto, latitud, longitud = campos
            longitud = longitud.rstrip("\n")
            g.agregar_vertice(codigo_aeropuerto)
            ciudades[ciudad] = ciudades.get(ciudad, set())
            ciudades[ciudad].add(codigo_aeropuerto)
            coordenadas[codigo_aeropuerto] = ((latitud, longitud), ciudad)


def cargarVuelos(ruta, g, dicVuelos):
    with open(ruta) as f:
        for linea in f:
            campos = linea.split(",")
            origen, destino, tiempo, precio, cantVuelos = campos
            cantVuelos = cantVuelos.rstrip("\n")
            g.agregar_arista(origen, destino, (int(tiempo), int(precio)))
            dicVuelos[origen] = dicVuelos.get(origen, {})
            dicVuelos[origen][destino] = int(cantVuelos)
            dicVuelos[origen][TOTAL] = dicVuelos[origen].get(TOTAL, 0) + int(cantVuelos)
            dicVuelos[destino] = dicVuelos.get(destino, {})
            dicVuelos[destino][origen] = int(cantVuelos)
            dicVuelos[destino][TOTAL] = dicVuelos[destino].get(TOTAL, 0) + int(cantVuelos)


def cargarItinerario(ruta):  # Devuelvo un grafo dirigido con precedencias
    with open(ruta) as f:
        ciudades = f.readline().split(",")
        ciudades[-1] = ciudades[-1].rstrip("\n")
        g = grafo(True, ciudades)
        restantes = f.readlines()
        for linea in restantes:
            campos = linea.split(",")
            campos[1] = campos[1].rstrip("\n")
            g.agregar_arista(campos[0], campos[1])
    return g


# Manejo de error
def perteneneCiudad(dic, origen, destino):
    if origen not in dic:
        print("La ciudad ", origen, " no tiene aeropuertos")
        return False
    if destino not in dic:
        print("La ciudad ", destino, " no tiene aeropuertos")
        return False
    return True


# Funciones de Grafo
def camino_minimo(g, origen, modo):  # modo[1] = barato / modo[0] = rapido
    dist = {}
    padre = {}
    for v in g:
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = []
    heapq.heappush(q, (0, origen))
    while len(q) > 0:
        peso, v = heapq.heappop(q)
        for w in g.adyacentes(v):
            if dist[v] + (g.peso(v, w))[modo] < dist[w]:
                dist[w] = dist[v] + (g.peso(v, w))[modo]
                padre[w] = v
                heapq.heappush(q, (dist[w], w))
    return dist, padre


def escalas_minimas(g, origen):  # BFS
    visitados = set()
    orden = {}
    padres = {}
    q = deque()
    visitados.add(origen)
    padres[origen] = None
    orden[origen] = 0
    q.append(origen)
    while q:
        v = q.popleft()
        for w in g.adyacentes(v):
            if w not in visitados:
                orden[w] = orden[v] + 1
                padres[w] = v
                visitados.add(w)
                q.append(w)
    return orden, padres


def minimos_generalizado(g, origen, destino, dic, modo, funcion):
    pesoMin, padresMin, destinoMin = float("inf"), [], str
    if modo == "rapido":
        var = 0
    else:
        var = 1
    for a in dic[origen]:
        for b in dic[destino]:
            if funcion == "camino_mas":
                peso, padres = camino_minimo(g, a, var)
            else:
                peso, padres = escalas_minimas(g, a)
            if peso[b] < pesoMin:
                pesoMin, padresMin, destinoMin = peso[b], padres, b
    return reconstruir_camino(padresMin, destinoMin)


def centralidad(g):
    cent = {}
    for v in g: cent[v] = 0
    for v in g:
        distancia, padre = escalas_minimas(g, v)
        cent_aux = {}
        for w in g: cent_aux[w] = 0
        vertices_ordenados = heapsort(distancia)
        for _, w in vertices_ordenados[::-1]:
            if padre[w] is None: continue
            cent_aux[padre[w]] += 1 + cent_aux[w]
        for w in g:
            if w == v: continue
            cent[w] += cent_aux[w]
    return cent


def calcular_grado_entrada(g):
    grado_entrada = {}
    for v in g:
        grado_entrada[v] = 0
    for v in g:
        for w in g.adyacentes(v):
            grado_entrada[w] += 1
    return grado_entrada


def topologia(g):
    grado = calcular_grado_entrada(g)
    res = []
    q = deque()
    for v in g:
        if grado[v] == 0:
            res.append(v)
            q.append(v)
    while q:
        v = q.pop()
        for w in g.adyacentes(v):
            grado[w] -= 1
            if grado[w] == 0:
                res.append(w)
                q.append(w)
    return res


def mst_prim(g):
    v = g.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    q = []
    arbol = grafo(False, g.obtener_vertices())
    for w in g.adyacentes(v):
        heapq.heappush(q, (g.peso(v, w)[1], (v, w)))  # precio, (origen,destino))
    while len(q) > 0:
        precio, tupla = heapq.heappop(q)
        v, w = tupla
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, g.peso(v, w))
        visitados.add(w)
        for u in g.adyacentes(w):
            if u not in visitados:
                heapq.heappush(q, (g.peso(w, u)[1], (w, u)))
    return arbol


# Funciones Auxiliares
def reconstruir_camino(padres, destino):
    res = [destino]
    padre = padres[destino]
    while padre is not None:
        res.append(padre)
        padre = padres[padre]
    ruta = " -> ".join(res[::-1])
    print(ruta)
    return res[::-1]


def mas_central(g, numero, cantVuelos):
    dicCentralidad = centralidad(g)
    mas_centrales = {}
    mas_vuelos = {}
    for elem in dicCentralidad:
        if dicCentralidad[elem] > 1:
            mas_centrales[elem] = cantVuelos[elem][TOTAL] + dicCentralidad[elem]
        else: mas_vuelos[elem] = cantVuelos[elem][TOTAL]
    centralesOrdenado = sorted(mas_centrales, key=mas_centrales.get, reverse=True)
    vuelosOrdenado = sorted(mas_vuelos, key=mas_vuelos.get, reverse=True)
    centralesOrdenado.extend(vuelosOrdenado)
    salida = centralesOrdenado[:numero]
    print(", ".join(salida))


def camino_minimo_por_orden(lista, g, dic):
    print(", ".join(lista))
    for i in range(len(lista) - 1):
        minimos_generalizado(g, lista[i], lista[i + 1], dic, 0, ESCALA)


def heapsort(iterable):
    h = []
    for clave, valor in iterable.items():
        if valor == float("inf"): continue
        heapq.heappush(h, (valor, clave))
    return [heapq.heappop(h) for _ in range(len(h))]


def exportar_KML(ruta, lista, coordenadas):
    with open(ruta, "w") as f:
        f.write("""<?xml version="1.0" encoding="UTF-8"?>\n""")
        f.write("""<kml xmlns="http://earth.google.com/kml/2.1">
    <Document>
        <name>KML de rutas</name>
        <description>Camino minimo en KML.</description>\n\n""")
        for elem in lista:  # elem = codigo_aeropuerto # COORDENADAS elem = ((lat,long),ciudad)
            coordenada, ciudad = coordenadas[elem]
            f.write(f"""        <Placemark>
            <name>{elem}</name>
            <description>Aeropuerto {elem} de {ciudad}</description>
            <Point>
                <coordinates>{coordenada[1]}, {coordenada[0]}</coordinates>
            </Point>
        </Placemark>\n\n""")

        for i in range(len(lista) - 1):
            coordenadaOrigen, _ = coordenadas[lista[i]]
            coordenadaDestino, _ = coordenadas[lista[i + 1]]
            f.write(f"""        <Placemark>
            <LineString>
                <coordinates>{coordenadaOrigen[1]}, {coordenadaOrigen[0]} {coordenadaDestino[1]}, {coordenadaDestino[0]}</coordinates>
            </LineString>
        </Placemark>\n\n""")
        f.write("""    </Document>
</kml>""")
    print("OK")


def nuevo_aeropuerto(g, ruta, dicVuelos):
    with open(ruta, "w") as f:
        writer = csv.writer(f)
        for origen, destino, tupla in g.obtener_aristas():
            tiempo, precio = tupla
            cantVuelos = dicVuelos[origen][destino]
            writer.writerow([origen, destino, tiempo, precio, cantVuelos])
    print("OK")
