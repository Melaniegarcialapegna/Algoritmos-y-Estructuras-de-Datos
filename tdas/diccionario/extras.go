// 	for nodoPadre != nil {

// 		// si queremos borrar el hijo IZQUIERDO
// 		if nodoPadre.izquierdo.clave == clave{
// 			//Si no tiene hijos
// 			if nodoPadre.izquierdo.izquierdo == nil && nodoPadre.izquierdo.derecho == nil {
// 				nodoPadre.izquierdo = nil
// 				abb.cantidad--
// 			}
			
// 			//Tiene solo un hijo
// 			if nodoPadre.izquierdo.izquierdo == nil && nodoPadre.izquierdo.derecho != nil{
// 				nodoPadre.izquierdo = nodoPadre.izquierdo.derecho
// 				abb.cantidad--
// 			}
// 			if nodoPadre.izquierdo.izquierdo != nil && nodoPadre.izquierdo.derecho == nil{
// 				nodoPadre.izquierdo = nodoPadre.izquierdo.izquierdo
// 				abb.cantidad--
// 			}

// 			//Tiene dos hijos
// 			if nodoPadre.izquierdo.izquierdo != nil && nodoPadre.izquierdo.derecho != nil{
// 				nodoReemplazo := nodoPadre.derecho.izquierdo // el mas chiquito de la derecha
// 				for nodoReemplazo.izquierdo != nil{
// 					nodoReemplazo = nodoReemplazo.izquierdo
// 				}
// 				nodoPadre.izquierdo.clave = nodoReemplazo.clave
// 				nodoPadre.izquierdo.dato = nodoReemplazo.dato
// 				abb.cantidad--
// 			}

// 			return nodoPadre.izquierdo.dato

// 			// si queremos borrar el hijo DERECHO
// 		} else if nodoPadre.derecho.clave == clave{
// 			//Si no tiene hijos
// 			if nodoPadre.derecho.izquierdo == nil && nodoPadre.derecho.derecho == nil {
// 				nodoPadre.derecho = nil
// 				abb.cantidad--
// 			}
			
// 			//Tiene solo un hijo
// 			if nodoPadre.derecho.izquierdo == nil && nodoPadre.derecho.derecho != nil{
// 				nodoPadre.derecho = nodoPadre.derecho.derecho
// 				abb.cantidad--
// 			}
// 			if nodoPadre.derecho.izquierdo != nil && nodoPadre.derecho.derecho == nil{
// 				nodoPadre.derecho = nodoPadre.derecho.izquierdo
// 				abb.cantidad--
// 			}

// 			//Tiene dos hijos
// 			if nodoPadre.derecho.izquierdo != nil && nodoPadre.derecho.derecho != nil{
// 				nodoReemplazo := nodoPadre.derecho.izquierdo // el mas chiquito de la derecha
// 				for nodoReemplazo.izquierdo != nil{
// 					nodoReemplazo = nodoReemplazo.izquierdo
// 				}
// 				nodoPadre.derecho.clave = nodoReemplazo.clave
// 				nodoPadre.derecho.dato = nodoReemplazo.dato
// 				abb.cantidad--
// 			}

// 			return nodoPadre.derecho.dato


// 		}else {//Si no es izquierdo ni el derecho
// 			condicion := abb.cmp(clave, nodoPadre.clave)
// 			//Si el nuevo es mas chico que el actual
// 			if condicion < 0 {
// 				nodoPadre = nodoPadre.izquierdo 


// 			} else if condicion > 0 { //Si el nuevo es mas grande que el actual
// 				nodoPadre = nodoPadre.derecho


//  //Si nos piden borrar la raiz
// 			} else if abb.raiz.clave == clave {
				
// 				//Si no tiene hijos
// 			if abb.raiz.izquierdo == nil && abb.raiz.derecho == nil {
// 				abb.raiz = nil
// 				abb.cantidad--
// 			}
			
// 			//Tiene solo un hijo
// 			if abb.raiz.izquierdo == nil && abb.raiz.derecho != nil{
// 				abb.raiz = nodoPadre.izquierdo
// 				abb.cantidad--
// 			}
// 			if abb.raiz.izquierdo != nil && abb.raiz.derecho == nil{
// 				abb.raiz = nodoPadre.derecho
// 				abb.cantidad--
// 			}

// 			//Tiene dos hijos
// 			if abb.raiz.izquierdo != nil && abb.raiz.derecho != nil{
// 				nodoReemplazo := nodoPadre.derecho.izquierdo // el mas chiquito de la derecha
// 				for nodoReemplazo.izquierdo != nil{
// 					nodoReemplazo = nodoReemplazo.izquierdo
// 				}
// 				nodoPadre.derecho.clave = nodoReemplazo.clave
// 				nodoPadre.derecho.dato = nodoReemplazo.dato
// 				abb.cantidad--
// 			}
// 			}
// 		}
// 	}