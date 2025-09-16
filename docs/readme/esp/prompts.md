# Prompts

Sería absurdo realizar todo el trabajo sin asistencia de IA hoy en día. Esto es lo que se utilizó.

## Asistencia de IA

Como asistencia de IA, utilicé el modelo Qwen 2.5 7b, que ejecuto en mi PC. No es pesado y puede ejecutarse fácilmente en cualquier PC moderno con al menos 16 GB de RAM, [leer](https://github.com/jzethar/Useful-Containers/blob/ollama/ollama/README.md).

Por qué autoalojado? Porque quiero tener control total sobre mis datos. No quiero depender de servicios de terceros que podrían no ser transparentes ni seguros. Además, es más rentable a largo plazo.

Básicamente, ayuda a autocompletar mi código y comentarios. También hay una opción para hacer preguntas y obtener respuestas.

Todo esto es posible con la extensión Continue en VSCode. Solo ejecuta el contenedor y conéctate a él en el puerto.

## Otras IAs utilizadas

Para este proyecto, también usé un grupo de IAs para verificar y revisar el código generado. Son:
- ChatGPT 5
- Grok
- Gemini

Principalmente se usaron ChatGPT 5 y Grok.

### Colaboración

Normalmente, la IA se puede usar para:
1. **Generación de código**: Defino previamente la interfaz que debo implementar, luego pido a la IA que genere el código.
2. **Pruebas**: Proporciono un fragmento de código y pido a la IA que escriba pruebas para él. Ayuda a automatizar el proceso de pruebas.
3. **Refactorización de código**: Proporciono un fragmento de código y pido a la IA que lo refactorice. En situaciones con fragmentos de código pequeños, no se usa tanto. Pero recuerdo una vez que el linter mostró un error por un código demasiado complejo. La IA me ayudó a dividirlo en partes más pequeñas.
4. **Traducción**: Viviendo una vida hablando cuatro idiomas todos los días, a veces necesito traducir. Puedo pedirle a la IA que traduzca usando el contexto.

## Configuración y Prompts

### Grok

Mi Grok siempre está configurado con este prompt:

```
La respuesta de Grok debe ser compacta sin profundizar a menos que se le pida.
```

El gran problema de Grok es que genera mucho texto innecesario que me hace perder tiempo leyéndolo. Este prompt lo mantiene conciso.

### ChatGPT 5

Para ChatGPT 5, no es necesario un prompt como este, ya que es lo suficientemente inteligente para dar respuestas cortas. Pero, si cambia a ChatGPT 4, se usará el mismo prompt.

Por ejemplo, para generar código uso esto:

```
¿Puedes generar una implementación en Go para la siguiente interfaz?
```

Lo mismo para las pruebas.

Normalmente genera excelentes respuestas y ahorra tiempo. Pero no siempre, así que debe haber un equilibrio entre generar y escribir por mí mismo.

## Notas Finales

La mejor solución es un asistente personal autoalojado. Si no es posible, usaré ChatGPT 5 o Grok como alternativa.