const JAR_EXTENSIONS = ['.jar', ]
const JAR_MIME_TYPES = ['application/java-archive', ]

export const isJarMime = (type) => {
    const typeLower = type?.toLowerCase() || ''
    return typeLower !== '' && JAR_MIME_TYPES.includes(typeLower)
}