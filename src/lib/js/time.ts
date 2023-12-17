export function Time(label?: string) {
    return function (target: unknown, name: string, descriptor: PropertyDescriptor) {
        const method = descriptor.value;

        descriptor.value = async function (...args: unknown[]) {
            const startTime = performance.now();
            const result = method.apply(this, args);
            const endTime = performance.now();

            const elapsed = endTime - startTime;
            console.log(`🕒 ${label ? label + ' ' : ''}${elapsed.toFixed(3)}ms`);

            return result;
        };
    };
}

export function sleep(millis: number) {
    const waitTill = new Date(new Date().getTime() + millis);
    while (waitTill > new Date()) {
    }
}
