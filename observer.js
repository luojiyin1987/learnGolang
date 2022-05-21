class Observer {
    constructor() {
        this.message = {};
    }

    on(type, fn) {
        if (!this.message[type]) {
            this.message[type] = [];
        }
        this.message[type].push(fn);
    }

    off(type, fn) {
        if (!this.message[type]) return

        if (!fn) {
            this.message[type] = undefined;
            return;
        }

        this.message[type] = this.message[type].filter(item => item !== fn);
    }

    emit(type) {
        if (!this.message[type]) return;
        this.message[type].forEach(item => item());
    }
}

const person1 = new Observer();

person1.on('eat', handlerA);
person1.on('红宝书', handlerB);
person1.on('红宝书', handlerC);

function handlerA() {
    console.log('handlerA');
}
function handlerB() {
    console.log('handlerB');
}

function handlerC() {
    console.log('handlerC');
}
person1.emit('eat');
person1.emit('红宝书');
